# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make diffeomorphisms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

diffeomorphisms.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > diffeomorphisms.tex

diffeomorphisms.pdf: ../../*.tex ../../trademark.pdf *.tex diffeomorphisms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode diffeomorphisms.tex
	../../latexrun diffeomorphisms.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: diffeomorphisms.pdf
	open diffeomorphisms.pdf

o:
	make open

terms: diffeomorphisms.tex
	bbk sheet -mode ts diffeomorphisms.tex

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
	rm -f diffeomorphisms.*
	make reset
	bbk sheet -mode mk > Makefile