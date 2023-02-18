# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topological_neighborhoods.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

topological_neighborhoods.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > topological_neighborhoods.tex

topological_neighborhoods.pdf: ../../*.tex ../../trademark.pdf *.tex topological_neighborhoods.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topological_neighborhoods.tex
	../../latexrun topological_neighborhoods.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topological_neighborhoods.pdf
	open topological_neighborhoods.pdf

o:
	make open

terms: topological_neighborhoods.tex
	bbk_sheet -mode ts topological_neighborhoods.tex

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
	rm -f topological_neighborhoods.*
	make reset
	bbk_sheet -mode mk > Makefile