# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_errors_linear_model.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

normal_errors_linear_model.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > normal_errors_linear_model.tex

normal_errors_linear_model.pdf: ../../*.tex ../../trademark.pdf *.tex normal_errors_linear_model.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_errors_linear_model.tex
	../../latexrun normal_errors_linear_model.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_errors_linear_model.pdf
	open normal_errors_linear_model.pdf

o:
	make open

terms: normal_errors_linear_model.tex
	bbk_sheet -mode ts normal_errors_linear_model.tex

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
	rm -f normal_errors_linear_model.*
	make reset
	bbk_sheet -mode mk > Makefile