# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make joint_probability_distributions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

joint_probability_distributions.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > joint_probability_distributions.tex

joint_probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex joint_probability_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode joint_probability_distributions.tex
	../../latexrun joint_probability_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: joint_probability_distributions.pdf
	open joint_probability_distributions.pdf

o:
	make open

terms: joint_probability_distributions.tex
	bbk_sheet -mode ts joint_probability_distributions.tex

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
	rm -f joint_probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile