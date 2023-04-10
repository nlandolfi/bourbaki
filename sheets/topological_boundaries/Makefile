# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_boundaries.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topological_boundaries.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > topological_boundaries.tex

topological_boundaries.pdf: ../../*.tex ../../trademark.pdf *.tex topological_boundaries.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_boundaries.tex
	../../latexrun topological_boundaries.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_boundaries.pdf
	open topological_boundaries.pdf

o:
	make open

terms: topological_boundaries.tex
	bbk sheet -mode ts topological_boundaries.tex

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
	rm -f topological_boundaries.*
	make reset
	bbk sheet -mode mk > Makefile