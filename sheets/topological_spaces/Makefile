# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make topological_spaces.pdf
	make graph.pdf

topological_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > topological_spaces.tex

topological_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex topological_spaces.tex
	# pdflatex --file-line-error -interaction=nonstopmode topological_spaces.tex
	../../latexrun topological_spaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_spaces.pdf
	open topological_spaces.pdf

o:
	make open

terms: topological_spaces.tex
	bbk_sheet -mode ts topological_spaces.tex

spell:
	aspell -c topological_spaces.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f topological_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile