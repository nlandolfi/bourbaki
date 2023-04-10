# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make joint_cumulative_distribution_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

joint_cumulative_distribution_functions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > joint_cumulative_distribution_functions.tex

joint_cumulative_distribution_functions.pdf: ../../*.tex ../../trademark.pdf *.tex joint_cumulative_distribution_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode joint_cumulative_distribution_functions.tex
	../../latexrun joint_cumulative_distribution_functions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: joint_cumulative_distribution_functions.pdf
	open joint_cumulative_distribution_functions.pdf

o:
	make open

terms: joint_cumulative_distribution_functions.tex
	bbk sheet -mode ts joint_cumulative_distribution_functions.tex

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
	bbk sheet -mode mk > Makefile