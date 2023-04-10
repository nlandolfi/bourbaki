# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make polyhedra.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

polyhedra.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > polyhedra.tex

polyhedra.pdf: ../../*.tex ../../trademark.pdf *.tex polyhedra.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode polyhedra.tex
	../../latexrun polyhedra.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: polyhedra.pdf
	open polyhedra.pdf

o:
	make open

terms: polyhedra.tex
	bbk sheet -mode ts polyhedra.tex

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
	rm -f polyhedra.*
	make reset
	bbk sheet -mode mk > Makefile