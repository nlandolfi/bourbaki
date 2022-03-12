# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make joint_cumulative_distribution_functions.pdf

joint_cumulative_distribution_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > joint_cumulative_distribution_functions.tex

joint_cumulative_distribution_functions.pdf: ../../*.tex ../../trademark.pdf *.tex joint_cumulative_distribution_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode joint_cumulative_distribution_functions.tex
	../../latexrun joint_cumulative_distribution_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: joint_cumulative_distribution_functions.pdf
	open joint_cumulative_distribution_functions.pdf

o:
	make open

terms: joint_cumulative_distribution_functions.tex
	bbk_sheet -mode ts joint_cumulative_distribution_functions.tex

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
	rm -f joint_cumulative_distribution_functions.*
	make reset
	bbk_sheet -mode mk > Makefile