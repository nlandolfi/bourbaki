# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make featurized_probabilistic_linear_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

featurized_probabilistic_linear_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > featurized_probabilistic_linear_models.tex

featurized_probabilistic_linear_models.pdf: ../../*.tex ../../trademark.pdf *.tex featurized_probabilistic_linear_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode featurized_probabilistic_linear_models.tex
	../../latexrun featurized_probabilistic_linear_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: featurized_probabilistic_linear_models.pdf
	open featurized_probabilistic_linear_models.pdf

o:
	make open

terms: featurized_probabilistic_linear_models.tex
	bbk_sheet -mode ts featurized_probabilistic_linear_models.tex

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
	rm -f featurized_probabilistic_linear_models.*
	make reset
	bbk_sheet -mode mk > Makefile