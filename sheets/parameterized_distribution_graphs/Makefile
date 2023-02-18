# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make parameterized_distribution_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

parameterized_distribution_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > parameterized_distribution_graphs.tex

parameterized_distribution_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex parameterized_distribution_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode parameterized_distribution_graphs.tex
	../../latexrun parameterized_distribution_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: parameterized_distribution_graphs.pdf
	open parameterized_distribution_graphs.pdf

o:
	make open

terms: parameterized_distribution_graphs.tex
	bbk_sheet -mode ts parameterized_distribution_graphs.tex

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
	rm -f parameterized_distribution_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile