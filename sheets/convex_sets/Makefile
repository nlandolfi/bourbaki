# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make convex_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

convex_sets.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > convex_sets.tex

convex_sets.pdf: ../../*.tex ../../trademark.pdf *.tex convex_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode convex_sets.tex
	../../latexrun convex_sets.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_sets.pdf
	open convex_sets.pdf

o:
	make open

terms: convex_sets.tex
	bbk sheet -mode ts convex_sets.tex

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
	rm -f convex_sets.*
	make reset
	bbk sheet -mode mk > Makefile