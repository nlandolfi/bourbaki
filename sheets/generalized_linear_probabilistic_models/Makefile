# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generalized_linear_probabilistic_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

generalized_linear_probabilistic_models.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > generalized_linear_probabilistic_models.tex

generalized_linear_probabilistic_models.pdf: ../../*.tex ../../trademark.pdf *.tex generalized_linear_probabilistic_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generalized_linear_probabilistic_models.tex
	../../latexrun generalized_linear_probabilistic_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generalized_linear_probabilistic_models.pdf
	open generalized_linear_probabilistic_models.pdf

o:
	make open

terms: generalized_linear_probabilistic_models.tex
	bbk_sheet -mode ts generalized_linear_probabilistic_models.tex

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
	rm -f generalized_linear_probabilistic_models.*
	make reset
	bbk_sheet -mode mk > Makefile