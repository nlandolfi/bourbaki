# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make isomorphisms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

isomorphisms.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > isomorphisms.tex

isomorphisms.pdf: ../../*.tex ../../trademark.pdf *.tex isomorphisms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode isomorphisms.tex
	../../latexrun isomorphisms.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: isomorphisms.pdf
	open isomorphisms.pdf

o:
	make open

terms: isomorphisms.tex
	bbk sheet -mode ts isomorphisms.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f isomorphisms.*
	make reset
	bbk sheet -mode mk > Makefile